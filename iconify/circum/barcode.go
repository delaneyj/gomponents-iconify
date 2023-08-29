package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Barcode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.066 4.065H3.648a1.732 1.732 0 0 0-.963.189a1.368 1.368 0 0 0-.619 1.226v4.585a.5.5 0 0 0 1 0v-4.28a1.794 1.794 0 0 1 .014-.518c.077-.236.319-.2.514-.2h4.472a.5.5 0 0 0 0-1Zm-6.003 9.872v4.418a1.733 1.733 0 0 0 .189.963a1.369 1.369 0 0 0 1.227.619h4.584a.5.5 0 0 0 0-1h-4.28a1.831 1.831 0 0 1-.518-.014c-.236-.077-.2-.319-.2-.514v-4.472a.5.5 0 0 0-1 0Zm13.871 5.998h4.418a1.732 1.732 0 0 0 .963-.189a1.368 1.368 0 0 0 .619-1.226v-4.585a.5.5 0 0 0-1 0v4.28a1.794 1.794 0 0 1-.014.518c-.077.236-.319.2-.514.2h-4.472a.5.5 0 0 0 0 1Zm6.003-9.872V5.645a1.733 1.733 0 0 0-.189-.963a1.369 1.369 0 0 0-1.227-.619h-4.584a.5.5 0 0 0 0 1h4.28a1.831 1.831 0 0 1 .518.014c.236.077.2.319.2.514v4.472a.5.5 0 0 0 1 0Z"/><rect width="1" height="8.709" x="10.999" y="7.643" fill="currentColor" rx=".5"/><rect width="1" height="8.709" x="14.249" y="7.643" fill="currentColor" rx=".5"/><rect width="1" height="8.709" x="16.499" y="7.643" fill="currentColor" rx=".5"/><rect width="1" height="8.709" x="6.499" y="7.643" fill="currentColor" rx=".5"/><rect width="1.5" height="8.709" x="8.499" y="7.643" fill="currentColor" rx=".75"/>`),
		g.Group(children),
	)
}