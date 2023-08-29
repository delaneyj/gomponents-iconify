package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceFrowningCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilFaceFrowningCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M12.75 20.5a7.75 7.75 0 1 0 0-15.5a7.75 7.75 0 0 0 0 15.5Zm0 1a8.75 8.75 0 1 0 0-17.5a8.75 8.75 0 0 0 0 17.5Z" clip-rule="evenodd"/><path d="M11.5 10.25a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Zm5.5 0a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Z"/><path fill-rule="evenodd" d="M15.343 16.79a.5.5 0 0 0 .819-.573l-.412.283l.412-.284l-.002-.002l-.002-.002l-.004-.007l-.013-.018a1.984 1.984 0 0 0-.182-.202a2.89 2.89 0 0 0-.56-.417C14.892 15.272 14.117 15 13 15s-1.89.272-2.4.568a2.89 2.89 0 0 0-.558.417a2 2 0 0 0-.183.202l-.013.018l-.004.007l-.002.002l-.001.002l.411.284l-.412-.283a.5.5 0 1 0 .83.56a.984.984 0 0 1 .07-.074c.07-.068.188-.168.365-.27c.35-.205.95-.433 1.897-.433c.946 0 1.547.228 1.897.432c.177.103.295.203.365.27a.986.986 0 0 1 .08.088Z" clip-rule="evenodd"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilFaceFrowningCircleFilled0)"/></g>`),
		g.Group(children),
	)
}