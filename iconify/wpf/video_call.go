package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoCall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M8.5 6C2.5 6 0 6.813 0 7.469V18.5c0 .641 2.5 1.5 8.5 1.5s8.5-.813 8.5-1.469V7.5C17 6.859 14.5 6 8.5 6zm16.875.031a.885.885 0 0 0-.469.157l-6.5 4.187c-.25.188-.406.498-.406.813v3.624c0 .315.156.624.406.813l6.5 4.188c.176.133 1.032.55 1.032-.813V7c0-.787-.271-.981-.563-.969z"/>`),
		g.Group(children),
	)
}