package corelibs

import (
	"testing"
	"time"
)

func TestGetBobyResponseTime(t *testing.T) {
	target := "https://www.example.com"

	metric := getBobyResponseTime(target)

	if metric.target != target {
		t.Errorf("target is %v, want %s", metric.target, target)
	}

	if metric.nsLookup == 0 {
		t.Errorf("nsLookup: %v", metric.nsLookup)
	}

	if metric.tcpConnection == 0 {
		t.Errorf("tcpConnection: %v", metric.tcpConnection)
	}

	if metric.tlsHandshake == 0 {
		t.Errorf("tlsHandshake: %v", metric.tlsHandshake)
	}

	if metric.serverProcessing == 0 {
		t.Errorf("serverProcessing: %v", metric.serverProcessing)
	}

	if metric.contentTransfer == 0 {
		t.Errorf("contentTransfer (%v)", metric.contentTransfer)
	}

	total := metric.nsLookup + metric.tcpConnection + metric.tlsHandshake + metric.serverProcessing + metric.contentTransfer
	if metric.total-total > (time.Duration(1) * time.Millisecond) {
		t.Errorf("total is %v, want %v (diff is too much: %v)", metric.total, total, metric.total-total)
	}

	if total > metric.total {
		t.Errorf("Sum of metric (%v) is > to total (%v)", total, metric.total)
	}

}
