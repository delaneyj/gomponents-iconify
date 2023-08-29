package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Door(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M59.188 59.188V29.189c0-36.533-54.375-35.97-54.375 0v29.998H2V62h60v-2.813h-2.812zm-4.721 0H9.535V31.675c0-32.989 44.932-33.505 44.932 0v27.513z"/><path fill="currentColor" d="M13.79 30.265h15.599v-18.5c-7.821 1.162-14.984 7.364-15.599 18.5m-.05 5.225h15.648v18.7H13.74zm20.873-23.754v18.529h11.411v5.225H34.613v18.7H50.26V43.12h3.265v-15H50.01c-1.346-9.944-8.082-15.381-15.397-16.384m12.349 17.79h5.625v12.188h-5.625V29.526"/><path fill="currentColor" d="M50.712 37.37a.938.938 0 1 0-1.492.753l-.392 2.505h1.892l-.392-2.505a.932.932 0 0 0 .384-.753m-.937-7.348a2.695 2.695 0 0 0-2.695 2.695a2.695 2.695 0 0 0 5.387 0a2.692 2.692 0 0 0-2.692-2.695m0 4.568a1.875 1.875 0 1 1-.002-3.75a1.875 1.875 0 0 1 .002 3.75"/>`),
		g.Group(children),
	)
}