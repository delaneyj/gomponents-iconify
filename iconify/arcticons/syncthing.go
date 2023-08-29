package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Syncthing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39 11.56A19.06 19.06 0 0 0 5.52 23.81m34.92 10.74A18.91 18.91 0 0 0 43.63 24a19.36 19.36 0 0 0-.31-3.43M8.05 33.48a19.07 19.07 0 0 0 24.83 7.67"/><circle cx="27.32" cy="25.62" r="5.06" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.44 11.35a5.06 5.06 0 1 1 0 10.11H40a5.06 5.06 0 0 1-1-9.89a4.8 4.8 0 0 1 1.44-.22Zm-4.32 20.86a5.06 5.06 0 1 1-1.41.2a5 5 0 0 1 1.41-.2ZM7.56 23.38a5.22 5.22 0 0 1 1.38.19a5.06 5.06 0 0 1-.87 9.9h-.51a5.05 5.05 0 0 1-2-9.68a4.94 4.94 0 0 1 2-.41Zm5.03 4.59l9.71-1.76m8.86-3.88l5.35-2.74m-5.97 9.94l2.8 3.51"/>`),
		g.Group(children),
	)
}