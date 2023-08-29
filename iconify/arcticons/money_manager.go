package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneyManager(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.072 20.726a12.476 12.476 0 0 1 3.647-5.67c-.932-.745 0-5.376 0-5.376s3.726 1.65 4.392 2.422c0 0 2.874-3.487 11.206-3.487S43.5 15.057 43.5 23.654s-6.229 11.339-6.229 11.339a19.977 19.977 0 0 1-1.49 4.392H32.48a21.385 21.385 0 0 1-1.41-2.529s-2.635.32-4.525.32a23.051 23.051 0 0 1-3.966-.426a6.403 6.403 0 0 1-.932 2.635h-3.593c-1.118-.958-1.677-3.966-1.677-3.966S7.327 31.32 5.41 28.87c-1.039-1.943-.906-6.255-.906-6.255a7.94 7.94 0 0 1 3.567-1.89Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.544 13.779a11.548 11.548 0 0 1 7.347-2.156a13.187 13.187 0 0 1 7.985 2.529"/><circle cx="11.612" cy="21.285" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}