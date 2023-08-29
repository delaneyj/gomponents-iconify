package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Magic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="m304.188 151l-29 93l-28-93l-93-29l93-28l28-94l29 94l93 28zm162-51l-14 47l-16-47l-46-15l46-14l16-46l14 46l46 14zm-457 511l573-556c4-5 10-8 17-8c9-1 18 2 25 9c3 2 7 6 13 11c11 10 23 22 32 34c4 6 6 11 6 16c1 10-2 19-8 25l-573 558c-5 5-11 7-18 8c-9 1-18-2-25-9c-3-3-7-6-13-12c-11-9-23-21-31-34c-4-6-6-12-7-17c-1-10 2-18 9-25zm460-404l42 43l133-129l-42-44zm-279 57l-13 46l-14-46l-46-15l46-14l14-47l13 47l48 14zm466-35l15 47l46 14l-46 15l-15 46l-15-46l-46-15l46-14z"/>`),
		g.Group(children),
	)
}