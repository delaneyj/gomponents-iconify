package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackGi0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackGi1)"><use href="#flagpackGi0"/><path fill="#F7FCFF" d="M0 0h32v24H0z"/><path fill="#C51918" fill-rule="evenodd" d="M0 16h32v8H0v-8Z" clip-rule="evenodd"/><path fill="#DB000B" fill-rule="evenodd" d="M17.389 5H13.61v4.615h-.944V6.846H8.889v2.77H7v.923h.944v4.615H7v.923h17v-.923h-.944v-4.615H24v-.924h-1.889V6.846h-3.778v2.77h-.944V5Z" clip-rule="evenodd"/><path fill="#272727" d="M9.833 8.59a.82.82 0 0 1 1.64 0V9.8h-1.64V8.59Zm9.445 0a.82.82 0 0 1 1.64 0V9.8h-1.64V8.59Z"/><path fill="#272727" fill-rule="evenodd" d="M9.567 13.47a1.086 1.086 0 0 1 2.173 0v1.868H9.567v-1.867Z" clip-rule="evenodd"/><path fill="#272727" d="M19.012 13.47a1.086 1.086 0 0 1 2.172 0v1.868h-2.172v-1.867Z"/><path fill="#272727" fill-rule="evenodd" d="M13.895 13.64a1.392 1.392 0 0 1 2.783 0v2.67h-2.784v-2.67Z" clip-rule="evenodd"/><path fill="#272727" d="M14.556 7.666a.82.82 0 1 1 1.64 0V9.8h-1.64V7.666Z"/><path fill="#DE1A23" d="M7 8.692h.944v.923H7z"/><path fill="#DE1A23" fill-rule="evenodd" d="M8.5 6h1v1h-1V6Zm3.818 0h1v1h-1V6Zm-.568 0h-1.5v1h1.5V6Zm6.34 0h1v1h-1V6Zm3.728 0h1v1h-1V6Zm-.614 0h-1.5v1h1.5V6ZM13.8 4.023h-1v1h1v-1Zm1.5 0h-1v1h1v-1Zm.5 0h1v1h-1v-1Zm2.5 0h-1v1h1v-1Z" clip-rule="evenodd"/><path fill="#DE1A23" d="M22.999 8.775h1v.839h-1z"/><path fill="#E8AA00" fill-rule="evenodd" d="m15.44 18.453l2.712-1.774l-2.894-1.353l-2.894 1.663l3.075 1.464Zm-1.06-1.611l.888-.595l.937.51l-.843.552l-.982-.467Zm-.904 4.165h1.473v1.107v-.374h-.915v.374h-.558v-1.107Z" clip-rule="evenodd"/><path fill="#E8AA00" fill-rule="evenodd" d="M14.816 18.114h1.25v4h-1.25v-4Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackGi1"><use href="#flagpackGi0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}