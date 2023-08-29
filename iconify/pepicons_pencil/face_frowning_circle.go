package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceFrowningCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M12.75 20.5a7.75 7.75 0 1 0 0-15.5a7.75 7.75 0 0 0 0 15.5Zm0 1a8.75 8.75 0 1 0 0-17.5a8.75 8.75 0 0 0 0 17.5Z" clip-rule="evenodd"/><path d="M11.5 10.25a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Zm5.5 0a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Z"/><path fill-rule="evenodd" d="M15.343 16.79a.5.5 0 0 0 .819-.573l-.412.283l.412-.284l-.002-.002l-.002-.002l-.004-.007l-.013-.018a1.984 1.984 0 0 0-.182-.202a2.89 2.89 0 0 0-.56-.417C14.892 15.272 14.117 15 13 15s-1.89.272-2.4.568a2.89 2.89 0 0 0-.558.417a2 2 0 0 0-.183.202l-.013.018l-.004.007l-.002.002l-.001.002l.411.284l-.412-.283a.5.5 0 1 0 .83.56a.984.984 0 0 1 .07-.074c.07-.068.188-.168.365-.27c.35-.205.95-.433 1.897-.433c.946 0 1.547.228 1.897.432c.177.103.295.203.365.27a.986.986 0 0 1 .08.088Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}