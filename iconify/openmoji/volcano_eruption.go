package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolcanoEruption(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="23.575" cy="14.817" r="4.558" fill="#9b9b9a"/><circle cx="31.9" cy="14.393" r="5.923" fill="#3f3f3f"/><circle cx="36.505" cy="35.688" r="3.724" fill="#9b9b9a"/><circle cx="51" cy="22.028" r="5.064" fill="#9b9b9a"/><circle cx="40.101" cy="10.971" r="4.391" fill="#9b9b9a"/><circle cx="40.571" cy="22.005" r="7.821" fill="#3f3f3f"/><path fill="#9b9b9a" d="M7.948 64.288L21.46 44.389l9.749-7.234l6.263 1.473l5.316-1.473l5.128 6.412l8.347 5.824l6.556 14.897H7.948z"/><path fill="#d0cfce" d="M7.948 64.288L21.46 44.389l3.917-2.908l1.654 3.531l-1.614 6.284l1.893 12.992H7.948z"/><path fill="#3f3f3f" d="M62.819 64.288H43.646l3.881-12.137l-3.116-2.505l3.505-6.079l8.347 5.824l6.556 14.897z"/><circle cx="35.835" cy="28.595" r="5.643" fill="#9b9b9a"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m8.406 63.629l13.054-19.24l9.749-7.234l6.263 1.473l5.316-1.473l5.128 6.412l8.347 5.824l6.356 14.388"/><path d="m25.377 41.481l1.654 3.531l-1.614 6.284l.916 7.313m21.583-15.042l-3.505 6.079l3.116 2.505l-1.673 5.173M32.952 23.75a5.637 5.637 0 0 0-.264 9.528m11.804-22.307a4.386 4.386 0 0 0-8.658-.997a5.912 5.912 0 0 0-9.116 1.554a4.549 4.549 0 0 0-7.7 3.289m37.047 7.211a5.064 5.064 0 0 0-7.891-4.202m-9.631 15.719a5.65 5.65 0 0 0 2.806-3.758a7.812 7.812 0 0 0 5.806-3.571m-14.283-5.598a7.823 7.823 0 0 1 14.626-2.247"/></g>`),
		g.Group(children),
	)
}