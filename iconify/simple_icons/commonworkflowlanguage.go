package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Commonworkflowlanguage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.905 0L8.571 5.4l.037.037l.096.096l3.586 3.395l-2.24 2.252h-.01l-1.576 1.586l3.737 3.766l-3.735 3.803l.126.139v.012L12.052 24l1.608-1.64l-1.98-2.034l3.737-3.79l-1.608-1.642l-.01.012l-2.13-2.129l3.867-3.866l-.017-.015l.016-.016l-3.641-3.524l3.64-3.694z"/>`),
		g.Group(children),
	)
}