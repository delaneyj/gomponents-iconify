package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WinkTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0zm3 4c.552 0 1 .672 1 1.5S11.552 7 11 7s-1-.672-1-1.5s.448-1.5 1-1.5zm-5.5.876c.932 0 1.594.349 1.594.895c0 .116.06.672-.003.775c-.232-.384-.856-.659-1.591-.659s-1.359.275-1.591.659c-.062-.103-.003-.659-.003-.775c0-.546.662-.895 1.594-.895zM7.818 13c-1.863 0-3.498-1.004-4.42-2.515c1.1.86 3.04 1.028 5.083.625c2.191-.433 3.892-1.43 4.507-2.759C12.65 10.975 10.464 13 7.818 13z"/>`),
		g.Group(children),
	)
}