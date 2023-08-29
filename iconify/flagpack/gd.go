package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackGd0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackGd2)"><use href="#flagpackGd0"/><path fill="#C51918" fill-rule="evenodd" d="M0 0h32v24H0V0Z" clip-rule="evenodd"/><path fill="#FECA00" fill-rule="evenodd" d="M6 6h20v12H6V6Z" clip-rule="evenodd"/><mask id="flagpackGd1" width="20" height="12" x="6" y="6" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M6 6h20v12H6V6Z" clip-rule="evenodd"/></mask><g fill-rule="evenodd" clip-rule="evenodd" mask="url(#flagpackGd1)"><path fill="#23875F" d="m6 6l10 6l-10 6V6Zm20 0l-10 6l10 6V6Z"/><path fill="#C51918" d="M16 16a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path fill="#FECA00" d="m15.93 13.313l-2.072 1.437l.662-2.46L13 10.717l2.059-.085l.87-2.433l.871 2.433h2.055l-1.516 1.656l.76 2.316l-2.17-1.292Z"/></g><path fill="#FECA00" fill-rule="evenodd" d="m8.777 4.442l1.131-.702l1.186.631l-.415-1.131l.829-.81h-1.124l-.476-1.188l-.476 1.189l-1.125.041l.831.768l-.361 1.202Zm6 0l1.131-.702l1.186.631l-.415-1.131l.829-.81h-1.124l-.476-1.188l-.476 1.189l-1.125.041l.831.768l-.361 1.202Zm7.131-.702l-1.131.702l.361-1.202l-.83-.768l1.125-.041l.475-1.189l.476 1.189h1.123l-.828.809l.415 1.131l-1.186-.63ZM8.776 22.167l1.131-.702l1.186.631l-.415-1.13l.829-.81h-1.124l-.476-1.189l-.475 1.189l-1.125.041l.83.768l-.361 1.202Zm7.132-.702l-1.132.702l.361-1.202l-.83-.768l1.124-.041l.477-1.189l.475 1.189h1.123l-.828.81l.415 1.13l-1.185-.63Zm4.868.702l1.131-.702l1.186.631l-.415-1.13l.828-.81h-1.123l-.476-1.189l-.476 1.189l-1.125.041l.831.768l-.361 1.202Z" clip-rule="evenodd"/><path fill="#C51918" fill-rule="evenodd" d="M8.635 12.772s.538.48.832.643c.1-.384-.125-1.252-.125-1.252c-.063-.208-.758-.58-.758-.58s-.195.655.05 1.189Z" clip-rule="evenodd"/><path fill="#FECA00" fill-rule="evenodd" d="M9.196 13.93s-1.383-.646-1.524-1.56c-.14-.913.158-2.37.158-2.37s2.246.51 2.399 1.606c.152 1.095-.355 1.783-.355 1.783s-.52-1.335-1.01-1.483c0 0 0 1.323.332 2.024Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackGd2"><use href="#flagpackGd0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}