package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoJupyter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26.077 3.588a1.69 1.69 0 1 1-1.76-1.585a1.67 1.67 0 0 1 1.76 1.585zM16.219 23.11c-4.487 0-8.43-1.61-10.469-3.988a11.162 11.162 0 0 0 20.938 0c-2.034 2.378-5.962 3.988-10.469 3.988zm0-15.463c4.487 0 8.43 1.61 10.469 3.988a11.162 11.162 0 0 0-20.938 0c2.04-2.382 5.963-3.988 10.47-3.988zm-6.176 20.09a2.108 2.108 0 1 1-.203-.797a2.128 2.128 0 0 1 .203.798zM6.26 7.107A1.226 1.226 0 1 1 7.452 5.83A1.242 1.242 0 0 1 6.26 7.106z"/>`),
		g.Group(children),
	)
}