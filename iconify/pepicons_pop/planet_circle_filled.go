package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlanetCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopPlanetCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M13 6a7 7 0 1 0 0 14a7 7 0 0 0 0-14Zm-9 7a9 9 0 1 1 18 0a9 9 0 0 1-18 0Z"/><path d="M19.806 7.93a1 1 0 0 1-.367 1.365c-.676.39-1.564.688-2.55.897a1 1 0 1 1-.415-1.956c.872-.185 1.537-.427 1.966-.674a1 1 0 0 1 1.366.367ZM6.392 8.1a1 1 0 0 1 1.342-.445c1.033.518 2.958.915 5.266.915a1 1 0 0 1 0 2c-2.462 0-4.739-.414-6.162-1.127A1 1 0 0 1 6.392 8.1Zm14.363 3.317a1 1 0 0 1-.458 1.338C18.699 13.538 15.91 14 12.867 14c-2.87 0-5.494-.41-7.116-1.102a1 1 0 1 1 .784-1.84c1.27.541 3.588.942 6.332.942c2.931 0 5.358-.457 6.55-1.04a1 1 0 0 1 1.338.457Zm0 3.429a1 1 0 0 1-.458 1.338c-.853.418-2.247.719-3.706.92a35.96 35.96 0 0 1-4.734.325a1 1 0 1 1 0-2a33.96 33.96 0 0 0 4.46-.306c1.445-.2 2.559-.47 3.1-.735a1 1 0 0 1 1.338.458Zm-15.532.169a1 1 0 0 1 1.312-.528c.439.187 1.02.363 1.723.512a1 1 0 1 1-.412 1.957c-.789-.167-1.502-.377-2.095-.63a1 1 0 0 1-.528-1.311Z"/><path d="M9.857 16.286a.429.429 0 1 0 0 .857a.429.429 0 0 0 0-.857Zm-2.428.428a2.429 2.429 0 1 1 4.857 0a2.429 2.429 0 0 1-4.857 0Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopPlanetCircleFilled0)"/></g>`),
		g.Group(children),
	)
}