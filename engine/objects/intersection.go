package object

import (
	"math"
	"sort"

	"golang.org/x/exp/slices"

	"github.com/gernivisser/raytracer/engine/ray"
	"github.com/gernivisser/raytracer/tuples"
)

type Intersection struct {
	T      float64
	Object Object
}

type Intersections []Intersection

func NewIntersections(intersections ...Intersection) Intersections {
	var intrs Intersections
	for _, inter := range intersections {
		// Insterts ordered posistion
		intrs = *intrs.Insert(inter)
	}
	return intrs
}

// helper function to insert into ordered list
func (a Intersections) Insert(b Intersection) *Intersections {
	i := sort.Search(len(a), func(i int) bool { return a[i].T >= b.T })

	if i == len(a) {
		// Insert at end is the easy case.
		a = append(a, b)
	} else {
		a = append(a[:i+1], a[i:]...)
		a[i] = b
	}

	return &a
}

func (i Intersections) Hit() Intersection {
	hit := Intersection{T: math.Inf(1), Object: nil}

	for _, inter := range i {
		if inter.T >= 0 {
			if inter.T < hit.T {
				hit = inter
			}
		}
	}

	return hit
}

// Interaction Pre Computing

type IntersectComp struct {
	T           float64
	Object      Object
	Position    tuples.Tuple
	Eyev        tuples.Tuple
	Normalv     tuples.Tuple
	Over_point  tuples.Tuple
	Under_point tuples.Tuple
	Reflect_v   tuples.Tuple
	N1          float64
	N2          float64
}

var EPSILON float64 = 0.0001

func (i *Intersection) Perpare_computation(r ray.Ray, xs *Intersections) *IntersectComp {
	var res IntersectComp

	res.T = i.T
	res.Object = i.Object
	res.Position = *r.Position(i.T)
	res.Eyev = *r.Direction.Negate()
	res.Normalv = *i.Object.NormalAt(&res.Position)
	if res.Normalv.Dot(res.Eyev) < 0 {
		res.Normalv = *res.Normalv.Negate()
	}
	res.Over_point = *res.Position.Add(res.Normalv.Multiply(EPSILON))
	res.Under_point = *res.Position.Sub(res.Normalv.Multiply(EPSILON))
	res.Reflect_v = *r.Direction.Reflect(res.Normalv)

	//Compute N1 and N2
	var container []Object

	for _, interaction := range *xs {
		if interaction == *i {
			// If container is empty ray is moving thru air thus refrac_Index = 1
			if len(container) == 0 {
				res.N1 = 1
				// Else N1 assumes value of parent
			} else {
				res.N1 = container[len(container)-1].GetMaterial().Refractive_Index
			}
		}

		// Check Interaction in container list
		// If so Interaction is leaving object
		// If not Interaction is entering thus should be added to container
		inList := false
		for counter, obj := range container {
			if obj == interaction.Object {
				container = slices.Delete(container, counter, counter+1)
				inList = true
			}
		}
		if !inList {
			container = append(container, interaction.Object)
		}

		if interaction == *i {
			// If container is empty ray is exiting into air thus refrac_Index = 1
			if len(container) == 0 {
				res.N2 = 1
				// Else N2 assumes value of parent
			} else {
				res.N2 = container[len(container)-1].GetMaterial().Refractive_Index
			}
		}

	}

	return &res
}
