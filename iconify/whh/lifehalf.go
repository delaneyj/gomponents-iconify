package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lifehalf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 288q0 58-35 154.5T895 629q-35 55-55.5 86.5t-53.5 79t-56 75.5t-53.5 61t-55.5 51t-53.5 30t-55.5 12t-55.5-12.5T403 981t-56-51.5t-54-61.5t-56-75.5t-53.5-78.5t-54.5-85q-62-95-95.5-188.5T0 288Q0 169 84.5 84.5T288 0q67 0 125 28.5t99 78.5q40-50 98.5-78.5T736 0q119 0 203.5 84.5T1024 288zM736 128q-41 0-85 28.5T576 224q-3 4-11.5 18t-14 22.5t-16.5 16t-22 7.5v608q23 0 46-11.5t41-28.5t41-48t39.5-56.5T723 684l45-70q59-87 96-176t32-150q0-66-47-113t-113-47z"/>`),
		g.Group(children),
	)
}