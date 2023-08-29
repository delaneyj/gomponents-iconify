package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToonBlast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.005 41.408a7.453 7.453 0 0 0-2.953-1.878c.477-.084 2.048-.028 2.104-.814s-2.3-.786-2.3-.786s4.994-5.387 4.994-8.614c1.347.084 2.049-.955.842-2.61c1.066.196 1.179-1.993-.28-2.61c-.758-3.452-5.304-5.08-5.304-5.08s-1.207-7.1-7.016-9.26s-13.638 1.824-14.564 2.947c.056-2.301-3.732-4.658-5.865-1.516s-.757 5.81 2.133 5.95a9.45 9.45 0 0 0-2.245 3.928s-3.929 3.059-4.069 5.472c-.04.686.396.982 1.04 1.054c-1.012.63-1.713 2.538.476 3.689a6.482 6.482 0 0 0 1.571 2.245s-2.712-.016-2.61.87c.084.73 1.32.56 1.32.56A9.199 9.199 0 0 0 4.5 36.977"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.151 19.452c.526-4-2.441-7.346-4.188-7.998a2.62 2.62 0 0 0-3.347 1.179c-1.389-1.2-3.914-1.874-5.282-.064s-2.968 4.799-.442 9.408"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.795 19.452c2.625-.287 10.123.126 10.713 4.314s-5.368 7.877-8.711 8.027c-5.156.232-11.641-3.207-11.641-5.606s4.63-6.188 9.64-6.735Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.895 28.657c.253 1.15 1.98 6.159 4.645 6.481s4.98-3.334 4.98-3.334"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.68 32.826a2.901 2.901 0 0 1 3.169.067c.645-.411 2.16.237 2.42 1.346m6.615-22.837c.299-.706.886-4.238-2.785-4.768c-2.988-.431-2.622 2.544-2.622 2.544m-3.769 14.329c0 1.026 3.55 4.77 6.132 4.77s5.135-1.894 5.135-5.977c0-2.413-4.532-3.171-6.286-2.54s-4.981 1.53-4.981 3.747Z"/><circle cx="23.273" cy="18.049" r="1.08" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="27.679" cy="17.88" r=".94" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.12 11.117a2.587 2.587 0 0 1 3.073-1.543m4.125-1.207a7.476 7.476 0 0 1 3.297 2.035"/>`),
		g.Group(children),
	)
}