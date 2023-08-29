package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChronoTrigger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.258 25.723a12.089 12.089 0 0 0 10.744 10.026m4.189-.33a15.172 15.172 0 0 0 2.067-.705l.215-4.724a8.738 8.738 0 0 1-1.384.714m-6.116.185a7.824 7.824 0 0 1-4.472-4.1m.037-6.213a7.917 7.917 0 0 1 7.202-4.503a7.823 7.823 0 0 1 4.667 1.515l-.017-4.727a11.62 11.62 0 0 0-5.215-1.221a12.087 12.087 0 0 0-11.841 9.67"/><circle cx="28.096" cy="23.661" r="1.879" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.096 25.54v6.061m-1.879-7.941h-6.373m5.09 15.253A15.362 15.362 0 0 1 11.96 25.78m.038-4.383A15.365 15.365 0 0 1 24.85 8.545l.106-.016m3.154 15.194v0"/><circle cx="28.2" cy="5.161" r="2.661" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="28.079" cy="42.839" r="2.661" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="8.188" cy="23.661" r="2.661" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12.447 23.661l7.531 1.491m-7.531-1.491l7.531-1.491m8.118 15.696l-1.572-6.584m1.572 6.584l1.571-6.584m-4.122-19.533l-.308-1.776m-2.382 2.463l-.568-1.766m-1.642 2.886l-.971-1.533m-2.542 4.98l-1.429-1.061m.332 3.096l-1.677-.6m.978 2.862l-1.593-.182"/><circle cx="14.563" cy="28.877" r=".37" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="17.375" cy="33.264" r=".37" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="21.611" cy="36.284" r=".37" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.743 15.069l-1.259-1.289M30.67 35.286l.484 1.895m-.404-24.986l.46-2.355"/>`),
		g.Group(children),
	)
}