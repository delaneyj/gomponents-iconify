package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartFillTwentyFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14 20.408c-.492.308-.903.546-1.192.709c-.153.086-.308.17-.463.252h-.002a.75.75 0 0 1-.686 0a16.709 16.709 0 0 1-.465-.252a31.147 31.147 0 0 1-4.803-3.34C3.8 15.572 1 12.331 1 8.513C1 5.052 3.829 2.5 6.736 2.5C9.03 2.5 10.881 3.726 12 5.605C13.12 3.726 14.97 2.5 17.264 2.5C20.17 2.5 23 5.052 23 8.514c0 3.818-2.801 7.06-5.389 9.262A31.146 31.146 0 0 1 14 20.408Z"/>`),
		g.Group(children),
	)
}