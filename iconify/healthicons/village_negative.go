package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VillageNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsVillageNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM17.555 6.168a1 1 0 0 0-1.11 0l-6 4A1 1 0 0 0 10 11v8a1 1 0 0 0 1 1h4v-7h4v7h4a1 1 0 0 0 1-1v-8a1 1 0 0 0-.445-.832l-6-4Zm16.941 5.964a1 1 0 0 0-.992 0l-7 4A1 1 0 0 0 26 17v10a1 1 0 0 0 1 1h5v-8h4v8h5a1 1 0 0 0 1-1V17a1 1 0 0 0-.504-.868l-7-4ZM14.553 24.106l-8 4A1 1 0 0 0 6 29v12a1 1 0 0 0 1 1h6v-9h4v9h6a1 1 0 0 0 1-1V29a1 1 0 0 0-.553-.894l-8-4a1 1 0 0 0-.894 0Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsVillageNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}