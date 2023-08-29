package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cigarette(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M928.365 512q0 36 9 83q1 5 16.5 35.5t27 73.5t11.5 96l-64-64q-8-65-32-96q-23-24-27.5-47.5t-4.5-80.5q0-80 37.5-136t90.5-56q15 0 32 7q-42 16-69 67.5t-27 117.5zm-192-307q0 38 9 89q55 61 55 167v147l-64-64q1-7 1-23t-.5-35.5t-.5-24.5q0-77-9-89q-55-62-55-167q0-85 37.5-145t90.5-60q14 0 32 7q-42 18-69 72.5t-27 125.5zm278 764l-45 45q-10 10-23.5 10t-22.5-10l-658-657q-9-9-9-22.5t9-23.5l46-46q10-9 23.5-9t22.5 9l658 658q9 9 9 22.5t-10 23.5zm-758-745l-32 32q-3 3-7.5 8.5t-7.5 9t-7 6.5t-8 2.5t-7-3.5l-177-178q-10-9-10-22.5t10-23.5l45-46q10-9 23.5-9t22.5 9l178 178q3 3 3.5 7t-2.5 8t-6.5 7t-9 7.5t-8.5 7.5z"/>`),
		g.Group(children),
	)
}