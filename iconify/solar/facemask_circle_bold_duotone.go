package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FacemaskCircleBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M21.998 12.193L22 12c0-5.523-4.477-10-10-10S2 6.477 2 12c0 .13.002.258.007.386l4.98 1.81l2.877-1.15a5.75 5.75 0 0 1 4.271 0L17 14.192l4.998-1.999Z" opacity=".5"/><path d="M10 10.5c0 .828-.448 1.5-1 1.5s-1-.672-1-1.5S8.448 9 9 9s1 .672 1 1.5Zm6 0c0 .828-.448 1.5-1 1.5s-1-.672-1-1.5s.448-1.5 1-1.5s1 .672 1 1.5Zm-.234 10.767l.444-5.776l-2.632-1.052a4.25 4.25 0 0 0-3.156 0L7.79 15.49l.444 5.776A9.972 9.972 0 0 0 12 22c1.332 0 2.603-.26 3.766-.733Zm6.058-7.389l-4.112 1.644l-.38 4.94a10.009 10.009 0 0 0 4.492-6.584ZM6.668 20.461L6.29 15.54l-4.077-1.482a10.01 10.01 0 0 0 4.456 6.404Z"/></g>`),
		g.Group(children),
	)
}