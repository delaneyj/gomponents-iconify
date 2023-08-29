package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xingse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.328 8.869c-.164.995.004 2.257-.616 2.888c-4.496 2.784-8.735 6.077-5.429 11.781c1.768 3.05 3.627 4.363 6.776 5.775c6.145 2.757 11.115-3.996 9.664-8.277c-.16-.472 1.45-2.963 2.965-3.581c-2.954-.873-4.585-.779-5.544-.154a34.983 34.983 0 0 0 .115-4.697a16.047 16.047 0 0 0-4.582 2.079c-1.138-1.908-.356-4.277-3.35-5.814Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.384 27.223c-2.46 3.638-4.955 7.269-1.907 11.908m1.849-5.929c2.616 2.437 5.391 3.13 8.316.77c-3.323-3.263-6.364-2.293-8.316-.77Zm-8.51-10.472c-.491 2.348-1.654 4.81 2.888 7.776c1.257-2.975.332-5.948-2.888-7.777ZM37.837 5.5c3.932 0 4.663 1.154 4.663 4.856m0 27.326c0 3.433-1.077 4.818-4.74 4.818m-27.288 0c-4.241 0-4.972-.461-4.972-4.933m0-27.172c0-3.933 1-4.895 5.087-4.895"/>`),
		g.Group(children),
	)
}