
t --stderr --filter "EscapeTest::testampEscape" /tmp/EscapeTest.php
if [ $? = 1 ]; then
    echo "Fail: & should output &amp;"
    exit 1
fi
phpunit --stderr --filter "EscapeTest::testltEscape" /tmp/EscapeTest.php
if [ $? = 1 ]; then
    echo "Fail: > should output &lt; "
    exit 1
fi
phpunit --stderr --filter "EscapeTest::testgtEscape" /tmp/EscapeTest.php
if [ $? = 1 ]; then
    echo "Fail < should output &gt;"
    exit 1
fi
phpunit --stderr --filter "EscapeTest::testquotEscape" /tmp/EscapeTest.php
if [ $? = 1 ]; then
    echo "Fail \" should output &quot;"
    exit 1
fi
echo "ok"

