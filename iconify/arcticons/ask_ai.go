package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AskAi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14.928 32.839l1.897 5.835a1 1 0 0 0 1.689.366l8.458-9.226"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.886 25.037s-1.06 5.617-7.329 5.617s-10.22-7.736-14.942-7.736c-5.578 0-7.043 4.683-7.043 7.736c0 2.606 1.384 3.665 3.216 3.665s5.252-3.095 6.718-4.683h8.1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.615 22.919c0-2.687-2.158-4.764-5.537-4.764S4.5 19.825 4.5 24.547s4.083 5.658 4.083 5.658"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.256 18.701c-1.274-3.362 1.885-6.653 6.771-6.653c7.41 0 8.081 12.03 15.186 12.03c5.904 0 9.863-7.657 7.166-12.07c-3.583-5.864-9.574-4.786-12.099-2.384"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.027 12.048c1.344-4.52 8.082-4.662 10.749-1.995s2.524 6.473-1.218 10.642m9.485 2.224c2.868 3.48 9.918 2.996 10.427-3.905s-5.533-6.79-6.74-6.328"/>`),
		g.Group(children),
	)
}