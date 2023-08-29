package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pccovid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.535 11.06c3.169 0 6.337 0 12.675-3.169c6.338 3.17 10.3 2.905 10.3 2.905"/><circle cx="7.95" cy="11.06" r="1.585" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="36.471" cy="11.06" r="1.585" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="22.21" cy="38.524" r="1.585" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M20.626 38.524a24.438 24.438 0 0 1-5.901-4.694m-1.858-2.294C9.81 27.244 7.95 21.652 7.95 14.757m28.52-2.112c0 .774.054 1.832.07 3.09M36.382 20c-.52 5.435-2.578 12.25-9.947 16.939"/><path d="M7.95 27.433c-4.753 4.225-3.697 6.866-2.113 6.866c2.258 0 4.528-.518 7.479-1.594C21.41 29.755 31.888 22.681 43.5 11.07m0 0l-1.057 3.169m-2.113-2.113l3.17-1.056m-21.29.613v14.573m0 4.574v3.47"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.21 22.03a3.17 3.17 0 1 1 0-6.121"/><circle cx="18.374" cy="13.746" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.805 16.226l-1.431-2.48"/><circle cx="16.166" cy="21.985" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.656 20.548l-2.49 1.438"/><circle cx="15.769" cy="17.463" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.33 18.15l-2.561-.687"/><circle cx="19.884" cy="24.592" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m20.57 22.031l-.686 2.56"/>`),
		g.Group(children),
	)
}