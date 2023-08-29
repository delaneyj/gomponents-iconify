package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bug(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1600 896q0 26-19 45t-45 19h-224q0 171-67 290l208 209q19 19 19 45t-19 45q-18 19-45 19t-45-19l-198-197q-5 5-15 13t-42 28.5t-65 36.5t-82 29t-97 13V576H736v896q-51 0-101.5-13.5t-87-33t-66-39T438 1354l-15-14l-183 207q-20 21-48 21q-24 0-43-16q-19-18-20.5-44.5T144 1461l202-227q-58-114-58-274H64q-26 0-45-19T0 896t19-45t45-19h224V538L115 365q-19-19-19-45t19-45t45-19t45 19l173 173h844l173-173q19-19 45-19t45 19t19 45t-19 45l-173 173v294h224q26 0 45 19t19 45zm-480-576H480q0-133 93.5-226.5T800 0t226.5 93.5T1120 320z"/>`),
		g.Group(children),
	)
}