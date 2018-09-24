package api

import (
	"errors"
	"fmt"
	"net"
)

type EgressDestinationsValidator struct{}

func (v *EgressDestinationsValidator) ValidateEgressDestinations(destinations []EgressDestination) error {
	if len(destinations) == 0 {
		return errors.New("missing destinations")
	}

	for _, destination := range destinations {

		if destination.Name == "" {
			return errors.New("missing destination name")
		}

		if len(destination.IPRanges) == 0 {
			return errors.New("missing destination IP range")
		}

		if destination.Protocol == "" {
			return errors.New("missing destination protocol")
		}

		if !isValidProtocol(destination.Protocol){
			return fmt.Errorf("invalid destination protocol '%s', specify either tcp, udp, or icmp", destination.Protocol)
		}

		if destination.Protocol != "icmp" && len(destination.Ports) == 0 {
			return errors.New("missing destination ports")
		}

		for _, portRange := range destination.Ports {
			if portRange.Start > portRange.End {
				return fmt.Errorf("invalid port range %d-%d, start must be less than or equal to end", portRange.Start, portRange.End)
			}

			if portRange.End > 65535 {
				return fmt.Errorf("invalid end port %d, must be in range 1-65535", portRange.End)
			}
		}

		if destination.Protocol != "icmp" && destination.ICMPCode != nil {
			return fmt.Errorf("invalid destination: cannot set icmp_code property for destination with protocol '%s'", destination.Protocol)
		}

		if destination.Protocol != "icmp" && destination.ICMPType != nil {
			return fmt.Errorf("invalid destination: cannot set icmp_type property for destination with protocol '%s'", destination.Protocol)
		}

		for _, ipRange := range destination.IPRanges {
			startIP := net.ParseIP(ipRange.Start)
			if startIP == nil || startIP.To4() == nil {
				return fmt.Errorf("invalid ip address '%s', must be a valid IPv4 address", ipRange.Start)
			}

			endIP := net.ParseIP(ipRange.End)
			if endIP == nil || endIP.To4() == nil {
				return fmt.Errorf("invalid ip address '%s', must be a valid IPv4 address", ipRange.End)
			}

			if ipRange.Start > ipRange.End {
				return fmt.Errorf("invalid IP range %s-%s, start must be less than or equal to end", ipRange.Start, ipRange.End)
			}
		}

	}
	return nil
}

func isValidProtocol(protocol string) bool {
	validProtocols := []string{"tcp", "udp", "icmp"}

	for _, validProtocol := range validProtocols {
		if validProtocol == protocol {
			return true
		}
	}
		return false
}