package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.19 4.156c-1.672-1.535-4.383-1.535-6.055 0L10 5.197L8.864 4.156c-1.672-1.535-4.382-1.535-6.054 0c-1.881 1.726-1.881 4.519 0 6.245L10 17l7.19-6.599c1.88-1.726 1.88-4.52 0-6.245zm-1.066 5.219L10 15.09L3.875 9.375c-.617-.567-.856-1.307-.856-2.094s.138-1.433.756-1.999c.545-.501 1.278-.777 2.063-.777c.784 0 1.517.476 2.062.978L10 7.308l2.099-1.826c.546-.502 1.278-.978 2.063-.978s1.518.276 2.063.777c.618.566.755 1.212.755 1.999s-.238 1.528-.856 2.095z"/>`),
		g.Group(children),
	)
}