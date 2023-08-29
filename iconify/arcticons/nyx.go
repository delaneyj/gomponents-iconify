package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nyx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.558 20.921a12.82 12.82 0 0 0-.575-3.034c-1.954-5.26-4.544-9.466-11.63-9.466c-7.574 0-15.525 4.923-15.525 14.497A56.627 56.627 0 0 0 13.019 42.5h23.585a53.147 53.147 0 0 0 2.434-15.264"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.781 18.962C11.091 9.435 17.067 5.5 24.694 5.5c6.275 0 12.864 2.272 16.034 14.613"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.44 18.56a4.132 4.132 0 0 0-4.074 4.412a4.826 4.826 0 0 0 3.618 4.732m7.416 2.138c.78 1.3 1.799 2.544 3.51 2.373a3.827 3.827 0 0 0 3.246-2.677m-13.998.225C8.692 31.06 5.311 33.98 7.61 35.955c1.094.94 2.428.556 3.525-.105"/><circle cx="15.751" cy="24.189" r="2.867" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="26.794" cy="24.189" r="2.867" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.349 19.817c1.817 0 2.285.83 2.285 3.318s-1.352 3.083-2.136 3.083s-2.434-.464-2.434-3.085s1.393-3.316 2.285-3.316Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.027 18.017a6.603 6.603 0 0 0-2.614 6.362a3.278 3.278 0 0 0 6.153 1.484M30.14 30.674c-.304 1.967 1.237 7.18 3.935 6.106s2.272-5.7 0-7.709"/>`),
		g.Group(children),
	)
}