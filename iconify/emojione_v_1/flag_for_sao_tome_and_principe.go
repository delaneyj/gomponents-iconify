package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForSaoTomeAndPrincipe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#3e9a00" d="M15.75 39.016L4.121 52.238C5.656 53.345 7.609 54 10 54h44c6.627 0 10-4.925 10-11v-4l-48.25.016zM54 10H10c-2.201 0-4.03.553-5.514 1.5L15.879 25h48.12v-4c0-6.075-3.373-11-10-11"/><path fill="#f9cb38" d="M1 23h63v18H1z"/><path fill="#ec1c24" d="M22.416 32L4.486 11.5C1.513 13.407 0 16.942 0 21v22c0 3.883 1.385 7.289 4.11 9.246l18.344-20.219l-.038-.027z"/><path fill="#25333a" d="m40.681 30.25l-4.371.01l-1.342-4.412l-1.346 4.412l-4.362-.01l3.536 2.68l-1.366 4.38l3.543-2.719l3.547 2.719l-1.37-4.38zm17.132 0l-4.363.01l-1.35-4.412l-1.345 4.412l-4.365-.01l3.539 2.68l-1.366 4.38l3.537-2.719l3.551 2.719l-1.371-4.38z"/>`),
		g.Group(children),
	)
}