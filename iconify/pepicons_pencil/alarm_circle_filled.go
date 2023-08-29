package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilAlarmCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M13 7.75a6 6 0 1 0 0 12a6 6 0 0 0 0-12Zm-7 6a7 7 0 1 1 14 0a7 7 0 0 1-14 0Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M6.146 20.604a.5.5 0 0 0 .708 0l1.5-1.5a.5.5 0 1 0-.708-.707l-1.5 1.5a.5.5 0 0 0 0 .707Zm11.5-2.208a.5.5 0 0 1 .708 0l1.5 1.5a.5.5 0 0 1-.708.708l-1.5-1.5a.5.5 0 0 1 0-.707ZM13 10.25a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-1 0v-3a.5.5 0 0 1 .5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M16.5 13.75a.5.5 0 0 1-.5.5h-3a.5.5 0 0 1 0-1h3a.5.5 0 0 1 .5.5Z" clip-rule="evenodd"/><path d="M5.381 5.546c2.023-1.931 3.951.376 3.951.376L5.574 9.51s-2.216-2.032-.193-3.964Z"/><path fill-rule="evenodd" d="M6.97 5.783c-.225.027-.53.135-.898.487c-.369.352-.49.651-.528.875c-.039.235.001.481.103.735c.017.043.036.084.055.125l2.12-2.025a2.148 2.148 0 0 0-.122-.06a1.364 1.364 0 0 0-.73-.137Zm1.602-.518c-.746-.494-1.95-.903-3.19.281c-1.24 1.185-.888 2.407-.43 3.175c.15.251.312.454.432.59c.113.128.19.2.19.2l3.758-3.59s-.067-.08-.19-.198a4.047 4.047 0 0 0-.57-.458Z" clip-rule="evenodd"/><path d="M20.619 5.546c-2.023-1.931-3.951.376-3.951.376l3.758 3.588s2.216-2.032.193-3.964Z"/><path fill-rule="evenodd" d="M19.03 5.783c.225.027.53.135.898.487c.369.352.49.651.528.875c.039.235-.001.481-.103.735a2.132 2.132 0 0 1-.055.125l-2.12-2.025c.04-.021.08-.042.122-.06c.249-.114.493-.165.73-.137Zm-1.602-.518c.746-.494 1.95-.903 3.19.281c1.24 1.185.888 2.407.43 3.175c-.15.251-.312.454-.432.59c-.113.128-.19.2-.19.2l-3.758-3.59s.067-.08.19-.198c.131-.127.326-.297.57-.458Z" clip-rule="evenodd"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilAlarmCircleFilled0)"/></g>`),
		g.Group(children),
	)
}