package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dollar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M263 148v256c65 47 107 89 131 125c24 37 35 73 36 110c2 60-44 175-167 195v119h-71V836C82 822 28 733 0 682l63-37c38 71 81 109 129 119V445c-56-43-93-75-107-94c-26-35-39-73-39-113c0-65 41-155 146-170V0h71v69c79 15 125 77 155 117l-60 45c-22-29-51-68-95-83zm-71 202V145c-54 14-70 63-70 90c0 46 29 78 47 95c3 4 12 11 23 20zm71 152v257c65-22 88-83 87-118c0-25-8-48-24-72c-14-19-35-42-63-67z"/>`),
		g.Group(children),
	)
}