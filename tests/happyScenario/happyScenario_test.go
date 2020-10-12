package happyScenario

import (
	"testing"
	"time"

	"github.com/spiral/endure"
	plugin12 "github.com/spiral/endure/tests/happyScenario/ProvidedValueButNeedPointer/plugin1"
	plugin22 "github.com/spiral/endure/tests/happyScenario/ProvidedValueButNeedPointer/plugin2"
	"github.com/spiral/endure/tests/happyScenario/plugin1"
	"github.com/spiral/endure/tests/happyScenario/plugin2"
	"github.com/spiral/endure/tests/happyScenario/plugin3"
	"github.com/spiral/endure/tests/happyScenario/plugin4"
	"github.com/spiral/endure/tests/happyScenario/plugin5"
	"github.com/spiral/endure/tests/happyScenario/plugin6"
	"github.com/spiral/endure/tests/happyScenario/plugin7"
	primitive "github.com/spiral/endure/tests/happyScenario/plugin8"
	"github.com/stretchr/testify/assert"
)

func TestEndure_DifferentLogLevels(t *testing.T) {
	testLog(t, endure.DebugLevel)
	testLog(t, endure.WarnLevel)
	testLog(t, endure.InfoLevel)
	testLog(t, endure.FatalLevel)
	testLog(t, endure.ErrorLevel)
	testLog(t, endure.DPanicLevel)
	testLog(t, endure.PanicLevel)
}

func testLog(t *testing.T, level endure.Level) {
	c, err := endure.NewContainer(level)
	assert.NoError(t, err)

	assert.NoError(t, c.Register(&plugin4.S4{}))
	assert.NoError(t, c.Register(&plugin2.S2{}))
	assert.NoError(t, c.Register(&plugin3.S3{}))
	assert.NoError(t, c.Register(&plugin1.S1{}))
	assert.NoError(t, c.Register(&plugin5.S5{}))
	assert.NoError(t, c.Register(&plugin6.S6Interface{}))
	assert.NoError(t, c.Init())

	res, err := c.Serve()
	assert.NoError(t, err)

	go func() {
		for r := range res {
			if r.Error.Err != nil {
				assert.NoError(t, r.Error.Err)
				return
			}
		}
	}()

	time.Sleep(time.Second * 2)
	assert.NoError(t, c.Stop())
}

func TestEndure_Init_OK(t *testing.T) {
	c, err := endure.NewContainer(endure.DebugLevel)
	assert.NoError(t, err)

	assert.NoError(t, c.Register(&plugin4.S4{}))
	assert.NoError(t, c.Register(&plugin2.S2{}))
	assert.NoError(t, c.Register(&plugin3.S3{}))
	assert.NoError(t, c.Register(&plugin1.S1{}))
	assert.NoError(t, c.Register(&plugin5.S5{}))
	assert.NoError(t, c.Register(&plugin6.S6Interface{}))
	assert.NoError(t, c.Init())

	res, err := c.Serve()
	assert.NoError(t, err)

	go func() {
		for r := range res {
			if r.Error.Err != nil {
				assert.NoError(t, r.Error.Err)
				return
			}
		}
	}()

	time.Sleep(time.Second * 2)
	assert.NoError(t, c.Stop())
}

func TestEndure_Init_1_Element(t *testing.T) {
	c, err := endure.NewContainer(endure.DebugLevel)
	assert.NoError(t, err)

	assert.NoError(t, c.Register(&plugin7.Plugin7{}))
	assert.NoError(t, c.Init())

	res, err := c.Serve()
	assert.NoError(t, err)

	go func() {
		for r := range res {
			if r.Error.Err != nil {
				assert.NoError(t, r.Error.Err)
				return
			}
		}
	}()

	time.Sleep(time.Second * 2)

	assert.NoError(t, c.Stop())
	time.Sleep(time.Second * 1)
}

func TestEndure_ProvidedValueButNeedPointer(t *testing.T) {
	c, err := endure.NewContainer(endure.DebugLevel)
	assert.NoError(t, err)

	assert.NoError(t, c.Register(&plugin12.Plugin1{}))
	assert.NoError(t, c.Register(&plugin22.Plugin2{}))
	assert.NoError(t, c.Init())

	res, err := c.Serve()
	assert.NoError(t, err)

	go func() {
		for r := range res {
			if r.Error.Err != nil {
				assert.NoError(t, r.Error.Err)
				return
			}
		}
	}()

	time.Sleep(time.Second * 2)

	assert.NoError(t, c.Stop())
	time.Sleep(time.Second * 1)
}

func TestEndure_PrimitiveTypes(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			println("test should panic")
		}
	}()
	c, err := endure.NewContainer(endure.DebugLevel)
	assert.NoError(t, err)

	assert.NoError(t, c.Register(&primitive.Plugin8{}))
	assert.NoError(t, c.Init())

	_, _ = c.Serve()
	assert.NoError(t, err)

	assert.NoError(t, c.Stop())
}