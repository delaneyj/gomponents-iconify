package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Female(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFemale0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M38.379 9.851c-5.468-5.467-14.332-5.467-19.8 0a13.956 13.956 0 0 0-4.1 9.9a13.96 13.96 0 0 0 4.1 9.9c5.468 5.467 14.332 5.467 19.8 0c5.467-5.468 5.467-14.332 0-19.8Z"/><path stroke-linecap="round" d="M18.464 29.535L5.736 42.263m13.435-.707L6.443 28.828"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFemale0)"/>`),
		g.Group(children),
	)
}