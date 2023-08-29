package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#40649F" fill-rule="nonzero"/><path fill="#FFF" d="M12.997 7.25c1.751 0 6.08.835 10.19 3.806c4.11 2.971 2.264 7.147 1.345 7.962c-.92.815-8.108 1.891-8.556 2.635c-.448.743-.97 3.597-3.21 3.597s-3.282-2.27-3.81-3.587c-.53-1.317-1.206-3.548-1.206-6.676S9.1 7.25 12.997 7.25zm-1.394 8.279l-.136.005c-1.353.101-1.99 2.085-1.847 4.247c.143 2.162 1.773 4.004 2.756 4.004c.983 0 2.44-1.306 2.15-4.13c-.291-2.825-1.706-4.223-3.059-4.121zm7.456-3.515c-1.895-.737-4.694-1.407-3.81.53c.298.651.538 1.783.721 3.395c3.524-.26 5.64-.68 6.345-1.26c1.057-.868-1.362-1.927-3.256-2.665z"/></g>`),
		g.Group(children),
	)
}