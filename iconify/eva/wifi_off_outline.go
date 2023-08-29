package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaWifiOffOutline0"><g id="evaWifiOffOutline1"><g id="evaWifiOffOutline2" fill="currentColor"><circle cx="12" cy="19" r="1"/><path d="m12.44 11l-1.9-1.89l-2.46-2.44l-1.55-1.55l-1.82-1.83a1 1 0 0 0-1.42 1.42l1.38 1.37l1.46 1.46l2.23 2.24l1.55 1.54l2.74 2.74l2.79 2.8l3.85 3.85a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Zm9.28-3.07A13.93 13.93 0 0 0 12 4a14.1 14.1 0 0 0-4.44.73l1.62 1.62a11.89 11.89 0 0 1 11.16 3a1 1 0 0 0 .69.28a1 1 0 0 0 .72-.31a1 1 0 0 0-.03-1.39ZM3.82 6.65a14.32 14.32 0 0 0-1.54 1.28a1 1 0 0 0 1.38 1.44a13.09 13.09 0 0 1 1.6-1.29ZM17 13.14a1 1 0 0 0 .71.3a1 1 0 0 0 .72-1.69A9 9 0 0 0 12 9h-.16l2.35 2.35A7 7 0 0 1 17 13.14Zm-9.57-2.88a8.8 8.8 0 0 0-1.9 1.49A1 1 0 0 0 7 13.14a7.3 7.3 0 0 1 2-1.41Zm1.1 5.14a1 1 0 1 0 1.39 1.44a3.06 3.06 0 0 1 3.84-.25l-2.52-2.52a5 5 0 0 0-2.71 1.33Z"/></g></g></g>`),
		g.Group(children),
	)
}