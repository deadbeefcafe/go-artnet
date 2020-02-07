package artnet

import (
	"errors"
	"net"
)

// Option is a functional option handler for Controller.
type Option func(*Controller) error

// SetOption runs a functional option against Controller.
func (c *Controller) SetOption(option Option) error {
	return option(c)
}

// MaxFPS sets the maximum amount of updates sent out per second
func MaxFPS(fps int) Option {
	return func(c *Controller) error {
		c.maxFPS = fps
		return nil
	}
}

// Set the polling broadcast address
func BcastAddr(addr string, port int) Option {
	return func(c *Controller) error {
		//c.maxFPS = fps
		ip := net.ParseIP(addr)
		if ip == nil {
			return errors.New("invalid ip address")
		}
		broadcastAddr = net.UDPAddr{
			//IP:   []byte{0xff, 0xff, 0xff, 0xff},
			IP: ip,
			//Port: int(packet.ArtNetPort),
			Port: port,
		}
		return nil
	}

}
