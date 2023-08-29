package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bluetooth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.649 15.938L11.829 12l6.82-3.94a.984.984 0 0 0 .5-.87a.968.968 0 0 0-.5-.861L12.029 2.5a.989.989 0 0 0-1 0a1 1 0 0 0-.5.87v7.769q-2.1-1.23-4.22-2.44c-.24-.139-.47-.279-.71-.409a.5.5 0 0 0-.51.86l4.95 2.85c-1.41.81-2.83 1.62-4.23 2.44c-.24.129-.48.27-.72.41a.5.5 0 0 0 .51.86c1.65-.951 3.28-1.891 4.93-2.85v7.769a.993.993 0 0 0 .5.871a.969.969 0 0 0 1 0l6.62-3.82a1.007 1.007 0 0 0 0-1.74Zm-7.12-12.57l6.62 3.82l-6.62 3.83Zm0 17.259v-7.639l6.62 3.82Z"/>`),
		g.Group(children),
	)
}