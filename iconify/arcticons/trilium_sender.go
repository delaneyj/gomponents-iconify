package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriliumSender(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.757 22.493c2.173 2.398 7.195 3.017 11.523-1.88c4.533-5.128 8.22-9.531 8.22-9.531S35.63 4.26 27.848 9.114s-5.597 11.717-4.091 13.379Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.757 22.493c1.773-3.804 5.009-8.307 12.398-11.062M21.858 22.493c1.12-2.796-.394-8.438-5.116-9.138s-9.619.962-12.242 4.853c3.28 1.224 4.94 4.986 10.843 6.4c2.667.64 5.64.069 6.515-2.115Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.858 22.493c-2.624-2.827-6.53-5.626-10.756-5.684m12.655 6.714c2.866 0 5.036 3.426 4.44 8.69a47.144 47.144 0 0 0-.196 8.416s-9.09-1.312-9.856-7.695c-.74-6.18 1.636-9.411 5.612-9.411Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.757 23.523c-1.9 2.722-2.49 8.974.156 12.515"/>`),
		g.Group(children),
	)
}