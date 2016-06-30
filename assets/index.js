var myChart = echarts.init(document.getElementById('main'));

// var ecConfig = require('echarts/config');
// var zrEvent = require('zrender/tool/event');
var curIndx = 0;
var mapType = [
    'china',
    // 23个省
    '广东', '青海', '四川', '海南', '陕西',
    '甘肃', '云南', '湖南', '湖北', '黑龙江',
    '贵州', '山东', '江西', '河南', '河北',
    '山西', '安徽', '福建', '浙江', '江苏',
    '吉林', '辽宁', '台湾',
    // 5个自治区
    '新疆', '广西', '宁夏', '内蒙古', '西藏',
    // 4个直辖市
    '北京', '天津', '上海', '重庆',
    // 2个特别行政区
    '香港', '澳门'
];

var convertData = function (data) {
    var res = [];
    //console.debug("geoMap:", geoCoordMap);
    for (var i = 0; i < data.length; i++) {
        var geoCoord = geoCoordMap[data[i].name];
        if (geoCoord) {
            res.push({
                name: data[i].name,
                value: geoCoord.concat(data[i].value)
            });
        }
    }
    return res;
};

option = {
    color: [
        'rgba(1, 1, 255, 1)',
        'rgba(1, 255, 0, 1)',
        'rgba(255, 0, 0, 1)',
    ],
    title : {
        text: '宜信全国门店帐号分布图',
        subtext: '中国',
        x:'center',
        // textStyle : {
        //     color: '#fff'
        // }
    },
    legend: {
        orient: 'vertical',
        x:'left',
        data:['帐号','门店'],
        textStyle : {
            color: '#222'
        }
    },
    toolbox: {
        show : true,
        orient : 'vertical',
        x: 'right',
        y: 'center',
        feature : {
            // mark : {show: true},
            // dataView : {show: true, readOnly: false},
            // restore : {show: true},
            saveAsImage : {show: true}
        }
    },
    // calculable : false,
    series : [
        {
            name: '帐号',
            type: 'map',
            mapType: 'china',
            selectedMode : 'single',
            data : [],
            itemStyle:{
                normal:{
                    label:{show:true},
                    borderColor:'rgba(100,149,237,0.8)',
                    borderWidth:1.5,
                },
                emphasis:{label:{show:true}}
            },
            markPoint : {
                symbolSize: 1,
                large: true,
                // effect : {
                //     show: true
                // },
                data: []
            }
        },
        {
            name: '门店',
            type: 'map',
            mapType: 'china',
            selectedMode : 'single',
            data : [],
            itemStyle:{
                normal:{
                    label:{show:true},
                    borderColor:'rgba(100,149,237,0.8)',
                    borderWidth:1.5,
                },
                // emphasis:{label:{show:true}}
                // emphasis: {
                //     borderColor: '#1e90ff',
                //     borderWidth: 5,
                //     label: {
                //         show: true
                //     }
                // }
            },
            markPoint : {
                symbol: 'diamond',
                symbolSize: 6,
                large: true,
                // effect : {
                //     show: true
                // },
                itemStyle:{
                    normal:{
                        label:{show:true},
                        borderColor:'rgba(100,149,237,0.8)',
                        borderWidth:1.5,
                    },
                    // emphasis:{label:{show:true}}
                    emphasis: {
                        borderColor: '#1e90ff',
                        // color: '#ffffff',
                        borderWidth: 6,
                        label: {
                            show: true
                        }
                    }
                },
                data: []
            }
        },
    ]
}

$.getJSON('/all-account', function(data) {
    console.debug("data size:", data.length);
    var res = [];
    for (var i = 0; i < data.length; i++) {
        res.push({
            name: data[i].Name,
            value: 1,
            geoCoord: [data[i].Lot, data[i].Lat]
        });
    }
    option.series[0].markPoint.data = res;
    myChart.setOption(option);
});

myChart.on(echarts.config.EVENT.MAP_SELECTED, function (param){
    console.debug("click event fired.");
    var len = mapType.length;
    var mt = mapType[curIndx % len];
    if (mt == 'china') {
        // 全国选择时指定到选中的省份
        var selected = param.selected;
        for (var i in selected) {
            if (selected[i]) {
                mt = i;
                while (len--) {
                    if (mapType[len] == mt) {
                        curIndx = len;
                    }
                }
                break;
            }
        }
        // option.tooltip.formatter = '点击返回全国<br/>{b}';
    }
    else {
        curIndx = 0;
        mt = 'china';
        // option.tooltip.formatter = '点击进入该省<br/>{b}';
    }
    option.series[0].mapType = mt;
    option.series[1].mapType = mt;
    option.title.subtext = mt + ' （点击返回）';
    myChart.setOption(option, true);
});
