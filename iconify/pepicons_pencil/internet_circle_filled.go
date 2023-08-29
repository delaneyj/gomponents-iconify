package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InternetCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilInternetCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M4.5 13a8.5 8.5 0 1 0 17 0a8.5 8.5 0 0 0-17 0Zm16 0a7.5 7.5 0 1 1-15 0a7.5 7.5 0 0 1 15 0Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M9.5 13c0 4.395 1.442 8 3.5 8s3.5-3.605 3.5-8c0-4.396-1.442-8-3.5-8s-3.5 3.604-3.5 8Zm6 0c0 3.889-1.245 7-2.5 7s-2.5-3.111-2.5-7s1.245-7 2.5-7s2.5 3.111 2.5 7Z" clip-rule="evenodd"/><path d="m6.735 8.312l.67-.742c.107.096.221.19.343.281c1.318.988 3.398 1.59 5.665 1.59c1.933 0 3.737-.437 5.055-1.19a5.59 5.59 0 0 0 .857-.597l.65.76c-.298.255-.636.49-1.01.704c-1.477.845-3.452 1.323-5.552 1.323c-2.47 0-4.762-.663-6.265-1.79a5.81 5.81 0 0 1-.413-.34Zm0 9.388l.67.74c.107-.096.221-.19.343-.28c1.318-.988 3.398-1.59 5.665-1.59c1.933 0 3.737.436 5.055 1.19c.321.184.608.384.857.596l.65-.76a6.583 6.583 0 0 0-1.01-.704c-1.477-.844-3.452-1.322-5.552-1.322c-2.47 0-4.762.663-6.265 1.789c-.146.11-.284.223-.413.34ZM5 13.5v-1h16v1H5Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilInternetCircleFilled0)"/></g>`),
		g.Group(children),
	)
}