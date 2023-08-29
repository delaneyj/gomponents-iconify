package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YandexLauncher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.144 17.207a9.756 9.756 0 0 0-8.245-3.366c-2.775.236-5.422-1.306-6.69-3.785a9.771 9.771 0 0 0-6.976-5.176c-4.673-.858-9.37 1.857-10.962 6.334a9.771 9.771 0 0 0 1.033 8.675c1.534 2.337 1.569 5.314.025 7.645a9.774 9.774 0 0 0-1.007 8.852c1.672 4.485 6.465 7.129 11.15 6.15a9.774 9.774 0 0 0 6.836-5.33c1.218-2.516 3.847-4.08 6.633-3.856a9.758 9.758 0 0 0 8.189-3.354c3.154-3.635 3.161-9.146.014-12.788Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.303 14.94a9.869 9.869 0 0 0-.574-3.187C37.138 7.276 32.44 4.562 27.767 5.42a9.849 9.849 0 0 0-3.37 1.282M23.604 41.3a9.812 9.812 0 0 0 3.925 1.776c4.684.979 9.477-1.665 11.149-6.15a9.795 9.795 0 0 0 .552-4.647M8.743 15.452a9.823 9.823 0 0 0-2.887 2.295c-3.147 3.642-3.14 9.153.014 12.789a9.827 9.827 0 0 0 2.828 2.254m20.573-14.477l-5.417 5.417l-5.418-5.417m5.418 5.417v6.097"/>`),
		g.Group(children),
	)
}