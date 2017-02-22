<?php
  require_once 'code.php';
  class EscapeTest extends PHPUnit_Framework_TestCase {

    public function testquotEscape() {
	$this->expectOutputString('&quot;');
	print h('"');
    }
    public function testgtEscape() {
	$this->expectOutputString('&gt;');
	print h('>');
    }
    public function testltEscape() {
	$this->expectOutputString('&lt;');
	print h('<');
    }
    public function testampEscape() {
	$this->expectOutputString('&amp;');
	print h('&');
    }

  }
?>

