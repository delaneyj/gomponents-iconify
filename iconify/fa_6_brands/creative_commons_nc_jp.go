package fa_6_brands

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreativeCommonsNcJp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M247.7 8C103.6 8 0 124.8 0 256c0 136.4 111.8 248 247.7 248C377.9 504 496 403.2 496 256C496 117.2 388.5 8 247.7 8zm.6 450.7c-112 0-203.6-92.5-203.6-202.7c0-21.1 3-41.2 9-60.3l127 56.5h-27.9v38.6h58.1l5.7 11.8v18.7h-63.8V360h63.8v56h61.7v-56h64.2v-35.7l81 36.1c-1.5 2.2-57.1 98.3-175.2 98.3zm87.6-137.3h-57.6v-18.7l2.9-5.6l54.7 24.3zm6.5-51.4v-17.8h-38.6l63-116H301l-43.4 96l-23-10.2l-39.6-85.7h-65.8l27.3 51l-81.9-36.5c27.8-44.1 82.6-98.1 173.7-98.1c112.8 0 203 90 203 203.4c0 21-2.7 40.6-7.9 59l-101-45.1z"/>`),
		g.Group(children),
	)
}