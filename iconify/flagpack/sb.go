package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackSb0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackSb2)"><use href="#flagpackSb0"/><path fill="#3D58DB" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackSb1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g mask="url(#flagpackSb1)"><path fill="#579D20" fill-rule="evenodd" d="M0 24h32V0L0 24Z" clip-rule="evenodd"/><path fill="#FECA00" d="m-.782 24.745l-1.804-2.496L33.675-1.954L35.48.542L-.781 24.745Z"/><path fill="#F7FCFF" fill-rule="evenodd" d="M3.653 6.547L5.017 5.6l1.428.85l-.5-1.524l.998-1.091H5.59l-.573-1.602l-.574 1.602l-1.355.056l1.001 1.035l-.436 1.62Zm5.658 0l1.364-.946l1.428.85l-.5-1.524l.999-1.091h-1.354l-.573-1.602l-.573 1.602l-1.356.056l1.001 1.035l-.436 1.62Z" clip-rule="evenodd"/><path fill="#F7FCFF" fill-rule="evenodd" d="m6.648 9.455l1.364-.946l1.428.85l-.5-1.524l.998-1.091H8.585l-.573-1.602l-.574 1.602l-1.355.056l1 1.035l-.435 1.62ZM5.017 11.6l-1.364.946l.436-1.62l-1.001-1.035l1.355-.056l.574-1.602l.573 1.602h1.353l-.998 1.09l.5 1.525l-1.428-.85Zm5.658 0l-1.364.946l.436-1.62l-1-1.035l1.355-.056l.573-1.602l.573 1.602h1.354l-1 1.09l.501 1.525l-1.428-.85Z" clip-rule="evenodd"/></g></g><defs><clipPath id="flagpackSb2"><use href="#flagpackSb0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}