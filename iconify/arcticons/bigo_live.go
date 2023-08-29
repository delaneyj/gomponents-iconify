package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BigoLive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M29.176 26.128c5.847-2.127 10.03-4.81 5.703-10.34c.297-11.146-18.474-13.271-23.182-.608c-.532 2.439-.36 7.699 1.487 11.921c.677 1.595.15 8.106-2.234 5.163c-1.123-1.522-2.157-1.585-2.6-1.422c-2.142 3.95 3.388 7.314 4.597 7.338c3.019 3.508 5.59 3.071 7.772 3.233"/><path d="M32.027 24.947c5.01 7.692 2.454 12.597-3.843 14.197M18.155 25.723c.978.913 2.396.734 3.383.055c.191-1.824 1.667-1.76 1.92.117c1.193-.024 2.44-.167 2.674.688c.136.915-.635 1.449-1.797 1.29c.054 1.834-.865 1.772-2.936 1.592c-.682.301-1.304.424-1.929.559m3.331-8.603c5.011 1.965 9.508 1.553 12.185-2.666"/></g><ellipse cx="25.085" cy="18.768" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.116" ry="1.277"/><ellipse cx="32.833" cy="16.883" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".93" ry="1.216"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.522 40.002c6.872 13.72 14.196-18.17.526-6.194m13.52-4.578c11.28-7.012 2.678 13.07-1.86 8.09M25.829 6.57c-4.112-3.578-6.752-2.194-8.554 1.676c-7.732-1.625-7.509 4.7-5.579 6.934c-3.365 2.433-4.098 4.866.124 7.299c-1.625.755-3.295 3.434 1.364 5.352c-1.182.949-2.77 1.831-.124 3.406m24.073-15.363h2.63m-1.315-1.506v2.925"/>`),
		g.Group(children),
	)
}