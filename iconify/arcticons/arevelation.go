package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Arevelation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.41 43.5h5.24a1.84 1.84 0 0 0 1.84-1.84h0V22a17.42 17.42 0 0 0-3.31-10.22l2.45-2.41a2.85 2.85 0 0 0 0-4h0a2.87 2.87 0 0 0-4 0l-2.46 2.41a17.41 17.41 0 0 0-20.34 0L11.4 5.36a2.88 2.88 0 0 0-4.06 0h0a2.85 2.85 0 0 0 0 4l2.44 2.41A17.39 17.39 0 0 0 6.51 22v19.66a1.84 1.84 0 0 0 1.84 1.84h20.935M16.64 24.78a3.22 3.22 0 1 1 3.22-3.22h0a3.22 3.22 0 0 1-3.22 3.21Zm14.72.01a3.22 3.22 0 1 1 3.22-3.22h0a3.22 3.22 0 0 1-3.22 3.21Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.25 31.776a4.84 4.84 0 0 1 9.575-.558l1.223.153l1.698-1.32l1.32 1.698l1.714-1.332l1.32 1.698l1.698-1.32l1.332 1.714l2.99.375l1.32 1.698l-1.698 1.32l-13.302-1.668a4.824 4.824 0 0 1-.9 1.299"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24.731 40.985l-3.747-2.07a4.832 4.832 0 1 1 1.476-2.658l1.079.596l2.066-.597l.596 2.067l2.086-.602l.596 2.066l2.067-.596l.601 2.086l2.638 1.456l.596 2.066l-2.066.597l-7.988-4.41"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.8 29.479c-2.204-.277-4.11.441-4.255 1.603s1.523 2.328 3.728 2.605a7.226 7.226 0 0 0 1.574.03"/>`),
		g.Group(children),
	)
}