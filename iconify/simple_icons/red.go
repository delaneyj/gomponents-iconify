package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Red(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 6.679V0L8.655 4.945Zm0 1.976v6.69l7.673-4L16.327 6.4zm0-1.976l3.345-1.734L12 0Zm8.655 6.133L12 17.322V24l12-6.242ZM12 24v-6.679l-8.655-4.509L0 17.758ZM4.327 11.345l7.673 4v-6.69L7.673 6.4Z"/>`),
		g.Group(children),
	)
}