package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Affiliatetheme(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M12.104 5c-1.867.016-4.728 1.95-7.067 4.955c-2.879 3.7-3.89 7.604-2.256 8.74c1.634 1.126 5.303-.962 8.182-4.66c2.879-3.699 3.89-7.602 2.256-8.728c-.307-.211-.685-.31-1.115-.307zm17.88 6c-3.8 6.359-9.896 9.542-13.625 7.266c-1.284-.786-2.074-2.451-2.365-4.266c-1.976 3.66-5.786 6.903-9.994 8c2.681 3.188 6.838 5 11.314 5c8.087 0 14.647-6.622 14.647-14.799c.01-.459.066-.764.023-1.201z"/>`),
		g.Group(children),
	)
}