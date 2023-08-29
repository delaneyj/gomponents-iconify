package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackSg0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackSg2)"><use href="#flagpackSg0"/><path fill="#F7FCFF" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackSg1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g fill-rule="evenodd" clip-rule="evenodd" mask="url(#flagpackSg1)"><path fill="#E31D1C" d="M0 0v12h32V0H0Z"/><path fill="#F1F9FF" d="M8.868 10.59S6.043 9.453 6.043 6.374s2.825-4.179 2.825-4.179c-1.373-.347-5.017-.037-5.017 4.179s3.589 5.01 5.017 4.216Zm.724-.51l.811-.489l.829.49l-.203-.955l.673-.754h-.912l-.387-.891l-.387.891l-.914.039l.675.715l-.185.955Zm3.757-.543l-.81.489l.185-.955l-.675-.715l.914-.038l.386-.892l.387.892h.912l-.673.753l.203.955l-.829-.49ZM11.287 5.77l.81-.49l.83.49l-.204-.955l.674-.753h-.912l-.387-.892l-.387.892l-.914.038l.675.715l-.185.955ZM9.37 7.223l-.81.489l.184-.955l-.675-.715l.914-.038l.387-.892l.387.892h.912l-.673.753l.202.955l-.828-.49Zm4.548.445l.81-.49l.829.49l-.203-.955l.673-.753h-.912l-.387-.892l-.386.892l-.915.039l.675.714l-.184.955Z"/></g></g><defs><clipPath id="flagpackSg2"><use href="#flagpackSg0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}