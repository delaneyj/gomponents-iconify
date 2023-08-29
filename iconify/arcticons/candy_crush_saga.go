package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CandyCrushSaga(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.392 25.966c5.18 4.664 6.616 2.98 9.929 6.308c2.738 2.75 2.773 6.442-.649 8.832s-10.48 1.92-17.179-3.093C6.853 32.294.783 22.45 9.267 17.623c6.741-3.836 9.2 4.81 13.125 8.343Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.292 16.742c.393-4.111 8.162-15.994 17.197-9.174c2.219 1.676 2.907 4.203 2.682 7.063"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.654 25.23c.081-4.092 2.26-11.65 12.491-10.384s9.562 13.429 5.941 17.12a12.551 12.551 0 0 1-5.983 3.2"/><circle cx="28.098" cy="24.044" r=".75" fill="currentColor"/><circle cx="28.516" cy="19.525" r=".75" fill="currentColor"/><circle cx="34.312" cy="21.418" r=".75" fill="currentColor"/><circle cx="33.211" cy="18.025" r=".75" fill="currentColor"/><circle cx="38.342" cy="20.275" r=".75" fill="currentColor"/><circle cx="37.831" cy="23.555" r=".75" fill="currentColor"/><circle cx="32.461" cy="24.579" r=".75" fill="currentColor"/><circle cx="40.108" cy="26.079" r=".75" fill="currentColor"/><circle cx="39.092" cy="30.331" r=".75" fill="currentColor"/><circle cx="37.081" cy="31.378" r=".75" fill="currentColor"/><circle cx="35.062" cy="26.829" r=".75" fill="currentColor"/><circle cx="33.822" cy="30.992" r=".75" fill="currentColor"/><circle cx="30.577" cy="28.147" r=".75" fill="currentColor"/><circle cx="24.002" cy="26.254" r=".75" fill="currentColor"/><circle cx="27.348" cy="15.109" r=".75" fill="currentColor"/><circle cx="25.502" cy="21.418" r=".75" fill="currentColor"/><circle cx="42.1" cy="20.506" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}