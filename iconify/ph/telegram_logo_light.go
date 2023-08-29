package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TelegramLogoLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M235.57 27.7a7 7 0 0 0-7.13-1.22L25.78 105.79a12.23 12.23 0 0 0 2.1 23.39L82 139.81V200a14 14 0 0 0 24.08 9.71l26.64-27.63l41.58 36.45a13.9 13.9 0 0 0 9.2 3.49a14.33 14.33 0 0 0 4.36-.69a13.86 13.86 0 0 0 9.34-10.17l40.62-176.59a7 7 0 0 0-2.25-6.87ZM30.05 117.37a.46.46 0 0 1 0-.32a.51.51 0 0 1 .15-.08l159.71-62.52l-103.3 74l-56.41-11.04Zm67.39 84A2 2 0 0 1 94 200v-51.89l29.69 26Zm88.07 7.08a1.93 1.93 0 0 1-1.34 1.44a2 2 0 0 1-2-.4l-84.53-74.15L223 45.5Z"/>`),
		g.Group(children),
	)
}