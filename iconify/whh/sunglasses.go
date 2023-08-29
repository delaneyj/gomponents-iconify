package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sunglasses(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1025 513q0 70-14.5 115.5T965 697t-68.5 31.5T801 737t-95.5-8.5T637 697t-45.5-68.5T577 513q0-26-19-45t-45.5-19t-45 19t-18.5 45q0 70-14.5 115.5T389 697t-68.5 31.5T225 737t-95.5-8.5T61 697t-45.5-68.5T1 513q0-57 8-89q-16-28-1-55L225 27q11-19 31.5-24.5t39 5.5t24 32.5T314 81L158 326q42-5 99-5q75 0 114.5 11t56.5 37q51 16 85 16t85-16q17-26 56.5-37T769 321q56 0 99 5L712 81q-11-19-5.5-40.5t24-32.5t39-5.5T801 27l217 342q15 27-1 55q8 32 8 89zM129 385q-30 0-47 18t-17 46t17 46t47 18q53 0 90.5-18.5T257 449t-37.5-45.5T129 385zm576 0q-30 0-47 18t-17 46t17 46t47 18q53 0 90.5-18.5T833 449t-37.5-45.5T705 385z"/>`),
		g.Group(children),
	)
}