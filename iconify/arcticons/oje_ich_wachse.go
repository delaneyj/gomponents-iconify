package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OjeIchWachse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M19.151 29.364c5.937 2.529 12.806.722 16.176-4.994c3.696-6.268 1.611-14.346-4.657-18.042c-6.268-3.696-14.346-1.61-18.042 4.658c-3.696 6.268-1.436 13.827 4.832 17.523"/><path d="M31.661 6.973c1.63 1.385 1.413 4.486-1.647 4.451s-2.644-2.88-1.178-3.06c1.442-.175 1.618 1.478.528 1.513m-8.931 10.674c1.011-1.907.87-3.963-.314-4.591s-2.965.41-3.976 2.317c-1.012 1.908-.815 3.814.37 4.442s2.909-.26 3.92-2.168Z"/><path d="M21.092 17.68a1.823 1.823 0 0 1-3.425-1.246m12.309 7.85c.782-2.012.403-4.037-.847-4.523s-2.897.752-3.68 2.764s-.364 3.883.885 4.369s2.86-.597 3.642-2.61Z"/><path d="M30.296 21.357a1.823 1.823 0 0 1-3.547-.84m-7.226 4.64c.973-.188 1.9.281 1.946 1.055"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.55 29.742c1.71-2.122 2.313-3.378 3.26-2.45c.537.527.08 1.313-1.292 2.778"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.412 31.148c-.105-1.266.71-1.955 1.44-1.786m.576 2.935c1.442.223 2.7-1.875 1.585-2.772"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.447 30.481c-1.36.046-3.282 2.344-1.864 3.61s2.727.27 3.268-.258s.753-1.005.87-1.522"/><ellipse cx="16.344" cy="39.665" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4.255" ry="2.796" transform="rotate(-54.89 16.344 39.665)"/><ellipse cx="23.618" cy="39.547" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.796" ry="4.255" transform="rotate(-47.264 23.618 39.547)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.63 37.353c-.379-.889-.133-2 .395-2.649m2.642 7.653c1.227 1.211 4.144 1.46 5.428.21m-6.244-8.734c1.726-.78 7.704-2.222 10.588 2.034c1.166 1.72.565 2.787-.565 3.45"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.83 30.282c2.403.75 5.421 5.856 3.088 6.903c-1.723.774-2.373-1.156-2.373-1.156c-.563-1.657-.932-2.343-1.372-2.905"/>`),
		g.Group(children),
	)
}