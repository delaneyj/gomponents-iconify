package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderSimpleDashedThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M125.6 83.2L95.73 60.8a4 4 0 0 0-2.4-.8H40a4 4 0 0 0-4 4v16a4 4 0 0 1-8 0V64a12 12 0 0 1 12-12h53.33a12.05 12.05 0 0 1 7.2 2.4l29.87 22.4a4 4 0 1 1-4.8 6.4ZM88 204H39.38a3.39 3.39 0 0 1-3.38-3.38V192a4 4 0 0 0-8 0v8.62A11.4 11.4 0 0 0 39.38 212H88a4 4 0 0 0 0-8Zm72 0h-32a4 4 0 0 0 0 8h32a4 4 0 0 0 0-8Zm64-56a4 4 0 0 0-4 4v48.89a3.12 3.12 0 0 1-3.11 3.11H200a4 4 0 0 0 0 8h16.89A11.12 11.12 0 0 0 228 200.89V152a4 4 0 0 0-4-4Zm-8-72h-48a4 4 0 0 0 0 8h48a4 4 0 0 1 4 4v24a4 4 0 0 0 8 0V88a12 12 0 0 0-12-12ZM32 156a4 4 0 0 0 4-4v-32a4 4 0 0 0-8 0v32a4 4 0 0 0 4 4Z"/>`),
		g.Group(children),
	)
}