package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rocket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.933 13.069s7.059-5.094 6.276-10.924a.465.465 0 0 0-.112-.268a.436.436 0 0 0-.263-.115C12.137.961 7.16 8.184 7.16 8.184c-4.318-.517-4.004.344-5.974 5.076c-.377.902.234 1.213.904.959l2.148-.811l2.59 2.648l-.793 2.199c-.248.686.055 1.311.938.926c4.624-2.016 5.466-1.694 4.96-6.112zm1.009-5.916a1.594 1.594 0 0 1 0-2.217a1.509 1.509 0 0 1 2.166 0a1.594 1.594 0 0 1 0 2.217a1.509 1.509 0 0 1-2.166 0z"/>`),
		g.Group(children),
	)
}