package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Odysee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.605 15.236c2.43.845 3.431-4.332 4.785-6.728c2-3.537 5.807-1.624 5.931 2.093c.132 3.945-3.777 8.62-7.177 11.563c-.917.794-.876 1.48.25 2.144c4.173 2.463 5.327 9.08 6.23 12.36c1.41 5.127-7.404 6.47-7.228.648c.074-2.414.59-5.464-4.087-6.33c1.423 10.76-5.408 11.49-10.666 11.514c-4.561.02-4.121-7.099.798-6.779c2.254.147 3.708-1.517 2.691-3.788a60.128 60.128 0 0 0-3.539-6.03c.102-.662-4.5 1.492-7.875 4.336c-3.319 2.796-5.906-2.82-2.292-6.48c2.214-2.242 8.572-3.987 8.572-3.987c-.29-4.051-4.318-8.622 2.293-12.959s11.303 3.006 11.314 8.423Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.98 8.259c-2.42.99-4.155 1.553-3.34 4.585c.714 2.654 2.646 3.89 5.732 2.442c2.12-.994 2.933-2.168 2.443-4.137c-.619-2.479-1.933-4.078-4.835-2.89Z"/>`),
		g.Group(children),
	)
}