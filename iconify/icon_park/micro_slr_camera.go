package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicroSlrCamera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path fill="#2F88FF" d="M44 26C44 28.6393 43.2696 31.1081 42 33.2152C40.8198 35.1737 39.1737 36.8198 37.2152 38C35.1081 39.2696 32.6393 40 30 40C27.3607 40 24.8919 39.2696 22.7848 38C18.7192 35.5502 16 31.0927 16 26C16 20.9073 18.7192 16.4498 22.7848 14C24.8919 12.7304 27.3607 12 30 12C32.6393 12 35.1081 12.7304 37.2152 14C39.1737 15.1802 40.8198 16.8263 42 18.7848C43.2696 20.8919 44 23.3607 44 26Z"/><path fill="#2F88FF" d="M4 14V38H22.7848C18.7192 35.5502 16 31.0927 16 26C16 20.9073 18.7192 16.4498 22.7848 14H4Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M37.2152 14C35.1081 12.7304 32.6393 12 30 12C27.3607 12 24.8919 12.7304 22.7848 14M37.2152 14H42V18.7848M37.2152 14C39.1737 15.1802 40.8198 16.8263 42 18.7848M22.7848 14H4V38H22.7848M22.7848 14C18.7192 16.4498 16 20.9073 16 26C16 31.0927 18.7192 35.5502 22.7848 38M37.2152 38C35.1081 39.2696 32.6393 40 30 40C27.3607 40 24.8919 39.2696 22.7848 38M37.2152 38H42V33.2152M37.2152 38C39.1737 36.8198 40.8198 35.1737 42 33.2152M42 18.7848C43.2696 20.8919 44 23.3607 44 26C44 28.6393 43.2696 31.1081 42 33.2152"/><rect width="9" height="5" x="8" y="9" fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M10 22V30"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M36 26C36 29.3137 33.3137 32 30 32M24 26C24 22.6863 26.6863 20 30 20"/></g>`),
		g.Group(children),
	)
}